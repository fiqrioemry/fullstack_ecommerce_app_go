import { productSchema } from "@/lib/schema";
import { productState } from "@/lib/constant";
import { FormInput } from "@/components/form/FormInput";
import { useProductMutation } from "@/hooks/useProduct";
import { SectionTitle } from "@/components/header/SectionTitle";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SwitchElement } from "@/components/input/SwitchElement";

const AddProduct = () => {
  const { createProduct } = useProductMutation();

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle
        title="Add New Product"
        description="Complete the product information to offer it to users."
      />
      <div className="bg-background rounded-xl shadow-sm border p-6">
        <FormInput
          className="w-full md:w-72"
          state={productState}
          schema={productSchema}
          text={"Add New Product"}
          isLoading={createProduct.isPending}
          action={createProduct.mutateAsync}
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
        </FormInput>
      </div>
    </section>
  );
};

export default AddProduct;
