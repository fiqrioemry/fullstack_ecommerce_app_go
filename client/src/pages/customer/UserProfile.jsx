import { formatDate } from "@/lib/utils";
import { Loading } from "@/components/ui/Loading";
import { useProfileQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { UploadAvatar } from "@/components/customer/profile/UploadAvatar";
import { UpdateProfile } from "@/components/customer/profile/UpdateProfile";

const UserProfile = () => {
  const { data, isLoading, isError, refetch } = useProfileQuery();

  if (isLoading) return <Loading className="mt-10" />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const { fullname, email, birthday, gender, phone, joinedAt } = data;

  return (
    <section className="min-h-[45vh]">
      <div className="bg-background rounded-xl shadow-lg p-6 flex flex-col md:flex-row items-center gap-8">
        {/* Avatar */}
        <div className="flex-shrink-0 space-y-8">
          <UploadAvatar profile={data} />
        </div>

        {/* Info */}
        <div className="w-full">
          <div className="flex items-center w-full md:w-1/2 justify-between pr-0 md:pr-2">
            <h2 className="text-2xl font-bold mb-2">{fullname}</h2>
            <UpdateProfile profile={data} edit="fullname" />
          </div>
          <p className="text-muted-foreground mb-4">{email}</p>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm text-muted-foreground">
            <div className="flex items-center  justify-between">
              <div>
                <span className="font-medium text-foreground">Birthday :</span>{" "}
                {birthday || "not set"}
              </div>
              <UpdateProfile profile={data} edit="birthday" />
            </div>
            <div className="flex items-center  justify-between">
              <div>
                <span className="font-medium text-foreground">Gender :</span>{" "}
                {gender || "not set"}
              </div>
              <UpdateProfile profile={data} edit="gender" />
            </div>
            <div className="flex items-center  justify-between">
              <div>
                <span className="font-medium text-foreground">Phone :</span>{" "}
                {phone || "not set"}
              </div>
              <UpdateProfile profile={data} edit="phone" />
            </div>
            <div>
              <span className="font-medium text-foreground">Joined At:</span>{" "}
              {formatDate(joinedAt)}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default UserProfile;
