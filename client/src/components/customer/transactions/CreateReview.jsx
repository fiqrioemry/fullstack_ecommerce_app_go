import { reviewSchema } from "@/lib/schema";
import { reviewState } from "@/lib/constant";
import { Button } from "@/components/ui/button";
import { useReviewMutation } from "@/hooks/useReview";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputRatingElement } from "@/components/input/InputRatingElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const CreateReview = ({ itemId }) => {
  const createReview = useReviewMutation();

  const handleCreateReview = (data) => {
    createReview.mutateAsync({ itemId, data });
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
      <InputFileElement name="image" label="Product Screenshot" isSingle />
      <InputTextareaElement
        name="comment"
        label="Comment"
        maxLength={200}
        placeholder="Enter comment for this product"
      />
    </FormAddDialog>
  );
};
