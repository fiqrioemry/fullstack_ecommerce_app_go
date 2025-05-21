import {
  Dialog,
  DialogTitle,
  DialogClose,
  DialogHeader,
  DialogContent,
  DialogTrigger,
  DialogDescription,
} from "@/components/ui/dialog";
import { EyeIcon } from "lucide-react";
import { formatDate } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { useCustomerDetail } from "@/hooks/useDashboard";

export const CustomerDetail = ({ userId }) => {
  const { data, isLoading, isError } = useCustomerDetail(userId);

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="icon" variant="outline">
          <EyeIcon />
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md">
        <DialogHeader>
          <DialogTitle className="text-center">User Details</DialogTitle>
          <DialogDescription className="text-center">
            Full information about the selected user.
          </DialogDescription>
        </DialogHeader>

        {isLoading ? (
          <div className="space-y-3">
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
          </div>
        ) : isError ? (
          <div className="text-red-500 text-sm">
            Failed to load user details.
          </div>
        ) : (
          <div className="space-y-4">
            <div className="flex items-center justify-center gap-4">
              <img
                src={data.avatar}
                alt={data.fullname}
                className="w-20 h-20 rounded-full object-cover border"
              />
            </div>

            <div className="text-sm space-y-2">
              <p>
                <span className="font-medium">Fullname :</span>{" "}
                {data.fullname || "-"}
              </p>
              <p>
                <span className="font-medium">Phone :</span> {data.phone || "-"}
              </p>
              <p>
                <span className="font-medium">Gender:</span>{" "}
                {data.gender || "-"}
              </p>
              <p>
                <span className="font-medium">Birthday :</span>{" "}
                {data.birthday || "-"}
              </p>
              <p>
                <span className="font-medium">Address :</span>{" "}
                {data.address || "-"}
              </p>

              <p>
                <span className="font-medium">Joined At:</span>{" "}
                {formatDate(data.createdAt)}
              </p>
              <p>
                <span className="font-medium">Last Login :</span>{" "}
                {data.lastLogin || "-"}
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
