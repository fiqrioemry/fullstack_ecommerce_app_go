import { reviewSchema } from "@/lib/schema";
import { reviewState } from "@/lib/constant";
import { Button } from "@/components/ui/button";
import { useOrderMutation } from "@/hooks/useOrder";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputRatingElement } from "@/components/input/InputRatingElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const CreateReview = ({ order }) => {
  const { createReview } = useOrderMutation();

  const handleCreateReview = (data) => {
    createReview.mutateAsync({ productId: order.ProductId, data });
  };

  return (
    <FormAddDialog
      state={reviewState}
      schema={reviewSchema}
      title="Create New Review"
      action={handleCreateReview}
      loading={createReview.isPending}
      buttonElement={
        <Button size="sm" type="button" className="w-28" variant="secondary">
          <span>Review</span>
        </Button>
      }
    >
      <InputRatingElement name="rating" label="Rating" />
      <InputTextareaElement
        name="comment"
        label="Comment"
        maxLength={200}
        placeholder="Enter comment for this product"
      />
    </FormAddDialog>
  );
};
