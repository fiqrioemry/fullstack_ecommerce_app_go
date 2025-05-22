import { Badge } from "@/components/ui/badge";
import { UpdateAddress } from "./UpdateAddress";
import { DeleteAddress } from "./DeleteAddress";
import { SetMainAddress } from "./SetMainAddress";
import { Card, CardContent } from "@/components/ui/card";

export const AddressCard = ({ address }) => {
  return (
    <Card
      className={
        address.isMain
          ? "border-blue-500 bg-blue-50"
          : "border border-border bg-card shadow-sm hover:shadow-md transition"
      }
    >
      <CardContent className="p-5 space-y-3">
        {/* Header */}
        <div className="flex justify-between items-center w-full">
          <div className="text-sm font-semibold text-foreground">
            {address.name}
          </div>
          {address.isMain && (
            <Badge variant="outline" className="text-xs">
              Utama
            </Badge>
          )}
        </div>

        {/* Detail */}
        <div className="text-sm text-start text-muted-foreground leading-snug w-full">
          <p>{address.phone || "-"}</p>
          <p>
            {address.address}, {address.subdistrict}, {address.district},{" "}
            {address.city}, {address.province} {address.postalCode}
          </p>
        </div>

        {/* Actions */}
        <div className="flex justify-start gap-2 pt-2 w-full">
          <UpdateAddress address={address} />
          <DeleteAddress address={address} />
          {!address.isMain && <SetMainAddress address={address} />}
        </div>
      </CardContent>
    </Card>
  );
};
