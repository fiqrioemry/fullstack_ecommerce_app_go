import {
  Dialog,
  DialogTitle,
  DialogClose,
  DialogContent,
  DialogTrigger,
  DialogDescription,
} from "@/components/ui/Dialog";
import { Button } from "@/components/ui/Button";
import { SubmitLoading } from "@/components/ui/SubmitLoading";

const FormToggle = ({
  title,
  onToggle,
  description,
  loading = false,
  buttonElement = (
    <Button size="sm" type="button">
      <span>Select</span>
    </Button>
  ),
}) => {
  return (
    <Dialog>
      <DialogTrigger asChild>{buttonElement}</DialogTrigger>
      <DialogContent className="sm:max-w-md rounded-xl p-6 space-y-6">
        {loading ? (
          <SubmitLoading text="Processing..." />
        ) : (
          <>
            <div className="text-center space-y-2">
              <DialogTitle className="text-2xl font-bold text-foreground">
                {title}
              </DialogTitle>
              <DialogDescription className="text-muted-foreground">
                {description}
              </DialogDescription>
            </div>

            <div className="flex justify-center gap-4 pt-4">
              <DialogClose asChild>
                <Button variant="secondary" className="w-32">
                  Cancel
                </Button>
              </DialogClose>

              <DialogClose asChild>
                <Button
                  variant="destructive"
                  className="w-32"
                  onClick={onToggle}
                >
                  Confirm
                </Button>
              </DialogClose>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};

export { FormToggle };
