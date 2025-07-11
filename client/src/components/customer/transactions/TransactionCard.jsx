import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card, CardContent } from "@/components/ui/Card";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { Link, useLocation, useNavigate } from "react-router-dom";

export const TransactionCard = ({ transactions }) => {
  const navigate = useNavigate();
  const location = useLocation();

  const openModal = (id) => {
    navigate(`/user/transactions/${id}`, {
      state: { backgroundLocation: location },
    });
  };
  return (
    <div className="space-y-6">
      {transactions.map((tx) => (
        <Card
          key={tx.id}
          className="border border-border bg-card shadow-sm hover:shadow-md transition"
        >
          <CardContent className="p-5 space-y-4">
            <div className="flex flex-col md:flex-row md:items-center md:justify-between w-full">
              <div className="space-y-1 text-start">
                <div className="text-sm text-muted-foreground">
                  Purchased – {formatDateTime(tx.createdAt)}
                </div>
                <div className="text-xs text-muted-foreground">
                  INV/{tx.id.slice(0, 8).toUpperCase()}
                </div>
              </div>
              <Badge
                variant={
                  tx.status === "success"
                    ? "success"
                    : tx.status === "pending"
                    ? "outline"
                    : "destructive"
                }
                className="capitalize w-fit text-xs mt-2 md:mt-0"
              >
                {tx.status}
              </Badge>
            </div>

            {/* Items */}
            <div className="border-t pt-4 flex gap-4 items-center w-full">
              <img
                src={tx.items[0]?.image}
                alt={tx.items[0]?.name}
                className="w-20 h-20 object-cover rounded border"
              />
              <div className="flex-1 text-start">
                <p className="font-medium text-sm text-foreground line-clamp-1">
                  {tx.items[0]?.name}
                </p>
                <p className="text-muted-foreground text-sm">
                  {tx.items.length} items
                </p>
              </div>
              <div className="text-right hidden md:block">
                <p className="text-sm text-muted-foreground">Total Amount</p>
                <p className="text-lg font-bold text-foreground">
                  {formatRupiah(tx.total)}
                </p>
              </div>
            </div>

            {/* Mobile total */}
            <div className="md:hidden text-right w-full">
              <p className="text-sm text-muted-foreground">Total Amount</p>
              <p className="text-lg font-bold text-foreground">
                {formatRupiah(tx.total)}
              </p>
            </div>

            {/* Actions */}
            <div className="pt-2 flex justify-end gap-3 w-full">
              {tx.status === "waiting_payment" && (
                <>
                  <Link to={tx.paymentLink}>
                    <Button size="sm" className="w-32" variant="secondary">
                      Payment link
                    </Button>
                  </Link>
                </>
              )}
              {(tx.status === "pending" ||
                tx.status === "success" ||
                tx.status === "process") && (
                <Button
                  size="sm"
                  className="w-32"
                  variant="secondary"
                  onClick={() => openModal(tx.id)}
                >
                  View Detail
                </Button>
              )}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
