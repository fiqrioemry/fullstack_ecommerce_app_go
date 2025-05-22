import { StarIcon } from "lucide-react";
import { formatDateTime } from "@/lib/utils";
import { useProductReviewsQuery } from "@/hooks/useReview";

const ReviewList = ({ product }) => {
  const { data: reviews = [] } = useProductReviewsQuery(product.id);

  if (!reviews || reviews.length === 0) return null;

  return (
    <div className="mt-10">
      <h3 className="text-xl font-semibold mb-4">Reviews ({reviews.length})</h3>
      <div className="space-y-6">
        {reviews.map((review) => (
          <div
            key={review.id}
            className="flex items-start gap-4 border border-gray-200 p-4 rounded-xl"
          >
            <img
              src={review.avatar}
              alt={review.fullname}
              className="w-12 h-12 rounded-full object-cover"
            />
            <div>
              <div className="flex items-center justify-between space-x-4 mb-1">
                <p className="font-medium text-gray-900">{review.fullname}</p>
                <span className="text-sm text-gray-500">
                  {formatDateTime(review.createdAt)}
                </span>
              </div>
              <div className="flex items-center gap-1 mb-1">
                {[...Array(5)].map((_, i) => (
                  <StarIcon
                    key={i}
                    size={16}
                    className={
                      i < review.rating ? "text-yellow-500" : "text-gray-300"
                    }
                    fill={i < review.rating ? "#facc15" : "none"}
                  />
                ))}
              </div>
              <p className="text-gray-700 text-sm">{review.comment}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export { ReviewList };
