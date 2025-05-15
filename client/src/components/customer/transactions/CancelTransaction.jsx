/* eslint-disable react/prop-types */
import { Button } from "@/components/ui/button";
import { useOrderMutation } from "@/hooks/useOrder";
import { FormToggle } from "@/components/form/FormToggle";

const CancelTransaction = ({ transaction }) => {
  const { cancelTransaction } = useOrderMutation();

  const handleCancelTransaction = () => {
    cancelTransaction.mutate(transaction.id);
  };
  return (
    <FormToggle
      title="Cancel Transaction"
      onToggle={handleCancelTransaction}
      loading={CancelTransaction.isPending}
      buttonElement={
        <Button size="sm" type="button" variant="destructive">
          <span>Cancel transaction</span>
        </Button>
      }
      description="Are you sure want to cancel this transaction ?"
    />
  );
};

export { CancelTransaction };
