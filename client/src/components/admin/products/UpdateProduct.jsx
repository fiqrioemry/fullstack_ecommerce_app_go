import { productSchema } from "@/lib/schema";
import { useProductMutation } from "@/hooks/useProduct";
import { SwitchElement } from "@/components/input/SwitchElement";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const UpdateProduct = ({ product }) => {
  const { updateProduct } = useProductMutation();

  return (
    <FormUpdateDialog
      state={product}
      title="Update Product"
      schema={productSchema}
      loading={updateProduct.isPending}
      action={updateProduct.mutateAsync}
    >
      <div className="space-y-4">
        <InputTextElement
          name="name"
          label="Product Name"
          placeholder="Enter the product name"
        />
        <SelectOptionsElement
          name="categoryId"
          data="category"
          label="Category"
          placeholder="Select category"
        />
        <InputTextareaElement
          maxLength={200}
          name="description"
          label="Product Description"
          placeholder="Enter the product description"
        />

        <div className="grid grid-cols-3 gap-4">
          <InputNumberElement
            name="price"
            label="Price (Rp)"
            placeholder="e.g. 500000"
          />
          <InputNumberElement
            name="stock"
            label="Stock"
            placeholder="e.g. 10"
          />
          <InputNumberElement
            name="discount"
            maxLength={2}
            label="Discount"
            placeholder="e.g. 10 (10%)"
          />
        </div>
        <div className="grid grid-cols-4 gap-4">
          <InputNumberElement
            name="weight"
            label="Weight (gr)"
            placeholder="min. 1000"
          />
          <InputNumberElement
            name="height"
            label="Height (cm)"
            placeholder="min. 10"
          />
          <InputNumberElement
            name="width"
            label="Width (cm)"
            placeholder="min. 10"
          />
          <InputNumberElement
            name="length"
            label="Length (cm)"
            placeholder="min. 10"
          />
        </div>
      </div>
      <InputFileElement
        name="images"
        label="Product Images"
        note="You can upload multiple images"
      />
      <SwitchElement
        name="isFeatured"
        label="Set as featured products (Display on homepage)"
      />
      <SwitchElement name="isActive" label="Set as active product" />
    </FormUpdateDialog>
  );
};
