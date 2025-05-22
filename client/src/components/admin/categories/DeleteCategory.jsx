import { FormDelete } from "@/components/form/FormDelete";
import { useCategoryMutation } from "@/hooks/useCategory";

export const DeleteCategory = ({ category }) => {
  const { deleteCategory } = useCategoryMutation();

  const handleDeleteCategory = () => {
    deleteCategory.mutateAsync(category.id);
  };
  return (
    <FormDelete
      title="Delete Category"
      onDelete={handleDeleteCategory}
      loading={deleteCategory.isPending}
      description="Are you sure want to delete this category?"
    />
  );
};
