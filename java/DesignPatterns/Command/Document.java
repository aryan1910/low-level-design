import java.util.*;

// Command interface
interface ActionListenerCommand{
    void execute();
}

// Receiver - performs the operation
class Document{
    public void open(){
        System.out.println("Opening document");
    }

    public void save(){
        System.out.println("Saving document");
    }
}

// Concrete command
class ActionOpen implements ActionListenerCommand  {
    private Document doc;

    public ActionOpen(Document doc){
        this.doc = doc;
    }
    
    @Override
    public void execute(){
        this.doc.open();
    }
}

// Concrete command
class ActionSave implements ActionListenerCommand{
    private Document doc;

    public ActionSave(Document doc){
        this.doc = doc;
    }
    
    @Override
    public void execute(){
        this.doc.save();
    }
}

// Invoker
class MenuOptions{
    private List<ActionListenerCommand> commands = new ArrayList<>();

    public void addCommand(ActionListenerCommand command){
        this.commands.add(command);
    }

    public void executeCommands(){
        for (ActionListenerCommand command: commands){
            command.execute();
        }
    }
}


class DocumentRunner{
    public static void main(String[] args) {
        Document doc = new Document(); // Receiver 

        ActionListenerCommand open = new ActionOpen(doc);
        ActionListenerCommand save = new ActionSave(doc);

        MenuOptions menuOptions = new MenuOptions();
        menuOptions.addCommand(save);
        menuOptions.addCommand(open);

        menuOptions.executeCommands();
    }
}
