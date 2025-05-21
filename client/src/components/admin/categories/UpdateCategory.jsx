import { categorySchema } from "@/lib/schema";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { useCategoryMutation } from "../../../hooks/useCategory";

export const UpdateCategory = ({ category }) => {
  const { updateCategory } = useCategoryMutation();

  return (
    <FormUpdateDialog
      state={category}
      title="Update Category"
      schema={categorySchema}
      loading={updateCategory.isPending}
      action={updateCategory.mutateAsync}
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
    </FormUpdateDialog>
  );
};
