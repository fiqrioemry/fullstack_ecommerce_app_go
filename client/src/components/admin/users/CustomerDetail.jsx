import {
  Dialog,
  DialogTitle,
  DialogClose,
  DialogContent,
  DialogHeader,
  DialogDescription,
} from "@/components/ui/Dialog";
import { formatDate } from "@/lib/utils";
import { Button } from "@/components/ui/Button";
import { Skeleton } from "@/components/ui/Skeleton";
import { useNavigate, useParams } from "react-router-dom";
import { useCustomerDetailQuery } from "@/hooks/useDashboard";

export const CustomerDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { data, isLoading } = useCustomerDetailQuery(id);

  return (
    <Dialog open={true} onOpenChange={() => navigate(-1)}>
      <DialogContent className="max-w-md">
        <DialogHeader>
          <DialogTitle>User Details</DialogTitle>
          <DialogDescription>
            Full information about the selected user.
          </DialogDescription>
        </DialogHeader>

        {isLoading ? (
          <div className="space-y-3">
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
          </div>
        ) : (
          <div className="space-y-4">
            <div className="flex items-center gap-4">
              <img
                src={data?.avatar}
                alt={data?.fullname}
                className="w-16 h-16 rounded-full object-cover border"
              />
              <div>
                <h3 className="text-lg font-semibold">{data?.fullname}</h3>
                <p className="text-sm text-muted-foreground">{data?.email}</p>
              </div>
            </div>
            <div className="text-sm space-y-2">
              <p>
                <span className="font-medium">Phone:</span> {data?.phone || "-"}
              </p>
              <p>
                <span className="font-medium">Gender:</span>{" "}
                {data?.gender || "-"}
              </p>
              <p>
                <span className="font-medium">Birthday:</span>{" "}
                {data?.birthday || "-"}
              </p>
              <p>
                <span className="font-medium">Bio:</span> {data?.bio || "-"}
              </p>
              <p>
                <span className="font-medium">Joined Since:</span>{" "}
                {formatDate(data?.createdAt)}
              </p>
            </div>
          </div>
        )}

        <DialogClose asChild>
          <Button variant="outline" className="w-full mt-4">
            Close
          </Button>
        </DialogClose>
      </DialogContent>
    </Dialog>
  );
};
