/* eslint-disable react/prop-types */
import { Button } from "@/components/ui/button";
import { useOrderMutation } from "@/hooks/useOrder";
import { FormToggle } from "@/components/form/FormToggle";

export const CancelOrder = ({ order }) => {
  const { createShipment } = useOrderMutation();

  const handleCancelOrder = () => {
    createShipment.mutate(order.id);
  };
  return (
    <FormToggle
      title="Cancel Order"
      onToggle={handleCancelOrder}
      loading={createShipment.isPending}
      buttonElement={
        <Button size="sm" type="button" variant="destructive">
          <span>Cancel Order</span>
        </Button>
      }
      description="Are you sure want to cancel this order ?"
    />
  );
};
