import { Skeleton } from "@/components/ui/skeleton";

export const ShipmentInfoSkeleton = () => {
  return (
    <div className="space-y-4">
      <Skeleton className="h-6 w-1/3" />

      <Skeleton className="h-6 w-24 rounded-full" />

      <div className="space-y-2">
        <Skeleton className="h-4 w-1/2" />
        <Skeleton className="h-4 w-2/3" />
        <Skeleton className="h-4 w-1/3" />
        <Skeleton className="h-4 w-1/4" />
      </div>

      <div className="space-y-1 mt-4">
        <Skeleton className="h-2 w-full rounded-full" />
        <div className="flex justify-between text-xs text-muted-foreground">
          <Skeleton className="h-3 w-10" />
          <Skeleton className="h-3 w-14" />
        </div>
      </div>
    </div>
  );
};
