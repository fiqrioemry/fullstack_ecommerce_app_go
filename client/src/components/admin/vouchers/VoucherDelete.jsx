import { useVoucherMutation } from "@/hooks/useVouchers";
import { FormDelete } from "@/components/form/FormDelete";

const VoucherDelete = ({ voucher }) => {
  const { deleteVoucher } = useVoucherMutation();

  const handleDeleteVoucher = () => {
    deleteVoucher.mutate(voucher.id);
  };

  return (
    <FormDelete
      title="Delete Vouchers"
      loading={deleteVoucher.isPending}
      onDelete={handleDeleteVoucher}
      description="Are you sure want to delete this Voucher?"
    />
  );
};

export { VoucherDelete };
