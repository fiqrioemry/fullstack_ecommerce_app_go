import { StarIcon } from "lucide-react";
import { formatDateTime } from "@/lib/utils";
import { useProductReviewsQuery } from "@/hooks/useReview";

export const ReviewList = ({ product }) => {
  const { data: reviews = [] } = useProductReviewsQuery(product.id);

  if (!reviews || reviews.length === 0) return null;

  return (
    <section className="mt-12">
      <h3 className="mb-6">Customer Reviews ({reviews.length})</h3>

      <div className="space-y-6">
        {reviews.map((review) => (
          <article
            key={review.id}
            className="flex items-start gap-4 p-5 bg-white border border-gray-200 rounded-2xl shadow-sm"
          >
            <img
              src={review.avatar}
              alt={review.fullname}
              className="w-14 h-14 rounded-full object-cover border border-gray-300"
            />
            <div className="flex-1 space-y-1">
              <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between">
                <p className="text-base font-semibold">{review.fullname}</p>
                <time className="text-sm text-muted">
                  {formatDateTime(review.createdAt)}
                </time>
              </div>

              <div className="flex items-center gap-0.5">
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

              <p className="text-sm text-gray-700 leading-relaxed mt-1">
                {review.comment}
              </p>
            </div>
          </article>
        ))}
      </div>
    </section>
  );
};
