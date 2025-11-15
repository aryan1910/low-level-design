package main

import (
	"context"
	"fmt"
)

type Resolver interface {
	Resolve(ctx context.Context, request string) bool
}

type ResolverNode struct {
	r    Resolver
	next *ResolverNode
}

func NewResolverNode(r Resolver) *ResolverNode {
	return &ResolverNode{
		r:    r,
		next: nil,
	}
}

func (c *ResolverNode) SetNext(r *ResolverNode) {
	c.next = r
}

func (r *ResolverNode) Resolve(ctx context.Context, request string) bool {
	if r.r.Resolve(ctx, request) {
		return true
	}
	if r.next != nil {
		return r.next.Resolve(ctx, request)
	}
	return false
}

type ResolverChain struct {
	head *ResolverNode
}

func NewResolverChain(p Resolver, sec []Resolver) *ResolverChain {
	head := NewResolverNode(p)

	temp := head
	for _, s := range sec {
		temp.next = NewResolverNode(s)
		temp = temp.next
	}

	return &ResolverChain{
		head: head,
	}
}

type HTTPResolver struct{}

func (h *HTTPResolver) Resolve(ctx context.Context, request string) bool {
	if request == "http" {
		fmt.Println("HTTPResolver handled the request.")
		return true
	}
	return false
}

type RedisResolver struct{}

func (r *RedisResolver) Resolve(ctx context.Context, request string) bool {
	if request == "redis" {
		fmt.Println("RedisResolver handled the request.")
		return true
	}
	return false
}

func main() {
	chain := NewResolverChain(&HTTPResolver{}, []Resolver{&RedisResolver{}})

	requests := []string{"http", "redis", "database"}
	for _, req := range requests {
		if chain.head.Resolve(context.Background(), req) {
			fmt.Printf("Request '%s' was resolved.\n", req)
		} else {
			fmt.Printf("Request '%s' was not resolved.\n", req)
		}
	}
}
