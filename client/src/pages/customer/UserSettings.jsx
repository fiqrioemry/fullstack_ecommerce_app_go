import {
  useNotificationSettingsQuery,
  useUpdateNotificationSetting,
} from "@/hooks/useNotification";
import { Loading } from "@/components/ui/Loading";
import { Switch } from "@/components/ui/Switch";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const groupedByTitle = (notifications) => {
  return notifications.reduce((acc, item) => {
    if (!acc[item.title]) acc[item.title] = [];
    acc[item.title].push(item);
    return acc;
  }, {});
};

const UserSettings = () => {
  const {
    data: notifications = [],
    isError,
    refetch,
    isLoading,
  } = useNotificationSettingsQuery();

  const { mutate: updateSetting } = useUpdateNotificationSetting();

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const grouped = groupedByTitle(notifications);

  return (
    <section className="section p-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Notification settings</h2>
        <p className="text-muted-foreground text-sm">
          Choose how you'd like to be notified
        </p>
      </div>

      {Object.entries(grouped).map(([title, list]) => (
        <div key={title} className="border-b pb-6 space-y-4">
          <div className="flex items-center justify-between">
            <h3 className="font-semibold text-lg">{title}</h3>
          </div>
          <div className="space-y-2 pl-4">
            {list.map((item) => (
              <div
                key={`${item.typeId}-${item.channel}`}
                className="flex items-center justify-between"
              >
                <span className="capitalize text-sm text-muted-foreground">
                  {item.channel}
                </span>
                <Switch
                  checked={item.enabled}
                  onCheckedChange={(val) =>
                    updateSetting({
                      typeId: item.typeId,
                      channel: item.channel,
                      enabled: val,
                    })
                  }
                  className="transition duration-200"
                />
              </div>
            ))}
          </div>
        </div>
      ))}
    </section>
  );
};

export default UserSettings;
