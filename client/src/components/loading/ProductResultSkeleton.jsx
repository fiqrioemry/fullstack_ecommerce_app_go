import { Skeleton } from "@/components/ui/Skeleton";

export const ProductResultsSkeleton = () => {
  return (
    <section className="section py-20 md:py-28 space-y-8">
      <div className="grid grid-cols-4 gap-4">
        <div className="col-span-4 md:col-span-1 space-y-6">
          <div className="space-y-2">
            <Skeleton className="h-6 w-1/2" />
            {[...Array(5)].map((_, i) => (
              <Skeleton key={i} className="h-4 w-3/4" />
            ))}
          </div>
          <div className="space-y-2 pt-6 border-t">
            <Skeleton className="h-6 w-1/2" />
            <div className="flex gap-2">
              <Skeleton className="h-10 w-full" />
              <Skeleton className="h-10 w-full" />
            </div>
            <Skeleton className="h-10 w-full" />
          </div>
          <div className="space-y-2 pt-6 border-t">
            <Skeleton className="h-6 w-1/2" />
            <Skeleton className="h-10 w-full" />
          </div>
        </div>

        {/* Main Content */}
        <div className="col-span-4 md:col-span-3">
          <div className="mb-4">
            <div className="flex items-center justify-between mb-4">
              <div className="flex gap-2">
                <Skeleton className="h-10 w-10" />
                <Skeleton className="h-10 w-10" />
              </div>
              <Skeleton className="h-10 w-44" />
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
              {[...Array(8)].map((_, i) => (
                <Skeleton key={i} className="h-64 w-full rounded-lg" />
              ))}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};
