import { shipmentSchema } from "@/lib/schema";
import { shipmentState } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { useOrderMutation } from "@/hooks/useOrder";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const ProceedOrder = ({ order }) => {
  const { createShipment } = useOrderMutation();

  const handleCreateShipment = (data) => {
    createShipment.mutateAsync({ orderId: order.id, data });
  };

  return (
    <FormAddDialog
      state={shipmentState}
      schema={shipmentSchema}
      title="Create New Shipment"
      action={handleCreateShipment}
      loading={createShipment.isPending}
      buttonElement={
        <Button size="sm" type="button" variant="secondary">
          <span>Proceed Shipment</span>
        </Button>
      }
    >
      <InputTextElement
        name="trackingCode"
        label="Tracking Number"
        placeholder="Enter the tracking number"
      />
      <InputTextareaElement
        name="note"
        label="Notes"
        maxLength={200}
        placeholder="Enter the notes for shipment"
      />
    </FormAddDialog>
  );
};
