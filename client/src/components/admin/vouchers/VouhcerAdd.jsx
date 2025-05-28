import { useNavigate } from "react-router-dom";
import { createVoucherSchema } from "@/lib/schema";
import { useVoucherMutation } from "@/hooks/useVouchers";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { SwitchElement } from "@/components/input/SwitchElement";
import { SelectElement } from "@/components/input/SelectElement";
import { createVoucherState, discountOptions } from "@/lib/constant";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const VoucherAdd = () => {
  const navigate = useNavigate();
  const { createVoucher } = useVoucherMutation();

  const handleCreateVoucher = async (data) => {
    const payload = {
      ...data,
      expiredAt: data.expiredAt ? new Date(data.expiredAt).toISOString() : null,
    };
    await createVoucher.mutateAsync(payload);
    navigate("/admin/vouchers");
  };

  return (
    <FormAddDialog
      title="Create New Voucher"
      state={createVoucherState}
      schema={createVoucherSchema}
      action={handleCreateVoucher}
      loading={createVoucher.isPending}
    >
      <InputTextElement name="code" label="Voucher Code " />
      <InputTextareaElement
        maxLength={200}
        name="description"
        label="Description"
        placeholder="e.g. Diskon 50% untuk semua kelas"
      />
      <SelectElement
        name="discountType"
        label="Discount Type"
        options={discountOptions}
        placeholder="Select discount type"
      />
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <InputNumberElement
          name="discount"
          label="Discount"
          placeholder="e.g. 50000 or 50"
        />
        <InputNumberElement
          name="maxDiscount"
          label="Max Discount (if %)"
          placeholder="e.g. 30000"
        />
        <InputNumberElement name="quota" label="Quota" placeholder="e.g. 10" />
      </div>
      <InputDateElement
        mode="future"
        name="expiredAt"
        label="Expiration Date"
      />
      <SwitchElement name="isReusable" label="Allow multiple usage?" />
    </FormAddDialog>
  );
};
