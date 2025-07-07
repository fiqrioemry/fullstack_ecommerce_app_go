import { PlusCircle } from "lucide-react";
import { addAddressSchema } from "@/lib/schema";
import { addressState } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { useAddressMutation } from "@/hooks/useAddress";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { LocationSelection } from "@/components/input/LocationSelection";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const AddAddress = () => {
  const { createAddress } = useAddressMutation();

  return (
    <FormAddDialog
      state={addressState}
      schema={addAddressSchema}
      title="Add New Address"
      action={createAddress.mutateAsync}
      loading={createAddress.isPending}
      buttonElement={
        <Button type="button">
          <PlusCircle className="w-4 h-4 mr-2" />
          <span>Address</span>
        </Button>
      }
    >
      <InputTextElement
        name="name"
        label="Nama penerima"
        placeholder="Masukkan nama penerima"
      />

      <InputTextElement
        name="phone"
        maxLength={12}
        isNumeric={true}
        label="Nomor Telepon"
        placeholder="Masukkan nomor telepon"
      />
      <InputTextareaElement
        name="address"
        label="Alamat Penerima"
        placeholder="Masukkan Alamat Penerima"
      />
      <LocationSelection />
    </FormAddDialog>
  );
};
