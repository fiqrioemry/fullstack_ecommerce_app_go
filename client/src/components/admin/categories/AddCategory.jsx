import { PlusIcon } from "lucide-react";
import { categorySchema } from "@/lib/schema";
import { categoryState } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { useCategoryMutation } from "@/hooks/useCategory";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputFileElement } from "@/components/input/InputFileElement";

export const AddCategory = () => {
  const { createCategory } = useCategoryMutation();

  return (
    <FormAddDialog
      state={categoryState}
      title="Create New Category"
      schema={categorySchema}
      loading={createCategory.isPending}
      action={createCategory.mutateAsync}
      buttonElement={
        <Button type="button">
          <PlusIcon />
          <span>Create Category</span>
        </Button>
      }
    >
      <div className="space-y-4">
        <InputTextElement
          name="name"
          label="Category Name"
          placeholder="Enter the category name"
        />
        <InputFileElement
          isSingle
          name="image"
          label="Category Image"
          placeholder="Upload category image"
          note="You can only upload .jpg, .jpeg, .png files"
        />
      </div>
    </FormAddDialog>
  );
};
