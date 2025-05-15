import { Button } from "@/components/ui/button";
import { useProfileQuery } from "@/hooks/useProfile";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { format } from "date-fns";
import { Camera } from "lucide-react";

const UserProfile = () => {
  const { data, isLoading, isError, error } = useProfileQuery();

  if (isLoading) return <Loading className="mt-10" />;
  if (isError)
    return <ErrorDialog message={error?.message || "Terjadi kesalahan"} />;
  if (!data) return null;

  const { fullname, email, birthday, gender, phone, avatar, joinedAt } = data;

  return (
    <section className="min-h-[45vh]">
      <div className="bg-background rounded-xl shadow-lg p-6 flex flex-col md:flex-row items-center gap-8">
        {/* Avatar */}
        <div className="flex-shrink-0 space-y-8">
          <img
            src={
              avatar ||
              `https://api.dicebear.com/6.x/initials/svg?seed=${fullname}`
            }
            alt={fullname}
            className="w-32 h-32 rounded-full border shadow"
          />
          <Button>
            <Camera /> Edit Avatar
          </Button>
        </div>

        {/* Info */}
        <div className="w-full">
          <h2 className="text-2xl font-bold mb-2">{fullname}</h2>
          <p className="text-muted-foreground mb-4">{email}</p>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm text-muted-foreground">
            <div>
              <span className="font-medium text-foreground">Birthday:</span>{" "}
              {birthday || "-"}
            </div>
            <div>
              <span className="font-medium text-foreground">Gender:</span>{" "}
              {gender || "-"}
            </div>
            <div>
              <span className="font-medium text-foreground">Phone:</span>{" "}
              {phone || "-"}
            </div>
            <div>
              <span className="font-medium text-foreground">Joined At:</span>{" "}
              {joinedAt ? format(new Date(joinedAt), "yyyy-MM-dd") : "-"}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default UserProfile;
