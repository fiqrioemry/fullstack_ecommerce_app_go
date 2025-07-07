import { Skeleton } from "@/components/ui/Skeleton";

export const TransactionDetailSkeleton = () => {
  return (
    <div className="space-y-4">
      <Skeleton className="h-6 w-1/3" />

      <div className="border p-4 rounded-md space-y-2 bg-muted">
        <Skeleton className="h-4 w-1/3" />
        <Skeleton className="h-4 w-1/2" />
        <Skeleton className="h-4 w-1/4" />
        <Skeleton className="h-8 w-28" />
      </div>

      <div className="border rounded-md divide-y">
        {[...Array(2)].map((_, i) => (
          <div key={i} className="flex gap-4 items-center p-4">
            <Skeleton className="w-16 h-16 rounded border" />
            <div className="flex-1 space-y-2">
              <Skeleton className="h-4 w-3/4" />
              <Skeleton className="h-4 w-1/2" />
            </div>
            <Skeleton className="h-8 w-28 rounded-md" />
          </div>
        ))}
      </div>

      <div className="border p-4 rounded-md space-y-3 bg-muted/50">
        <Skeleton className="h-5 w-1/4" />
        <Skeleton className="h-4 w-1/2" />
        <Skeleton className="h-4 w-3/4" />
        <Skeleton className="h-4 w-2/3" />
      </div>

      <div className="border p-4 rounded-md space-y-2">
        <Skeleton className="h-5 w-1/4" />
        <Skeleton className="h-4 w-1/2" />
        <Skeleton className="h-4 w-1/3" />
        <Skeleton className="h-4 w-1/4" />
        <Skeleton className="h-5 w-1/2" />
      </div>
    </div>
  );
};
