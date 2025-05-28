import { Skeleton } from "@/components/ui/skeleton";

export const ReviewListSkeleton = () => {
  return (
    <>
      {Array.from({ length: 3 }).map((_, idx) => (
        <article
          key={idx}
          className="flex items-start gap-4 p-5 bg-white border border-gray-200 rounded-2xl shadow-sm"
        >
          <Skeleton className="w-14 h-14 rounded-full" />
          <div className="flex-1 space-y-2">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2">
              <Skeleton className="w-32 h-4 rounded" />
              <Skeleton className="w-24 h-4 rounded" />
            </div>
            <div className="flex items-center gap-1">
              {Array.from({ length: 5 }).map((_, i) => (
                <Skeleton key={i} className="w-4 h-4 rounded" />
              ))}
            </div>
            <Skeleton className="w-full h-4 rounded" />
            <Skeleton className="w-5/6 h-4 rounded" />
          </div>
        </article>
      ))}
    </>
  );
};
