import { FormDelete } from "@/components/form/FormDelete";
import { useProductMutation } from "@/hooks/useProduct";

export const DeleteProduct = ({ product }) => {
  const { deleteProduct } = useProductMutation();

  const handleDeleteProduct = () => {
    deleteProduct.mutate(product.id);
  };
  return (
    <FormDelete
      title="Delete Product"
      loading={deleteProduct.isPending}
      onDelete={handleDeleteProduct}
      description="Are you sure want to delete this product ?"
    />
  );
};
