// src/components/address/AddAddress.jsx
import { addressSchema } from "@/lib/schema";
import { addressState } from "@/lib/constant";
import { useAddressMutation } from "@/hooks/useAddress";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { SwitchElement } from "@/components/input/SwitchElement";
import { PlusCircle } from "lucide-react";
import { Button } from "@/components/ui/button";
import LocationSelection from "@/components/input/LocationSelection";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const AddAddress = () => {
  const { createAddress } = useAddressMutation();

  return (
    <FormAddDialog
      state={addressState}
      schema={addressSchema}
      title="Add New Address"
      action={createAddress.mutateAsync}
      loading={createAddress.isPending}
      buttonElement={
        <Button size="sm" type="button">
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
      <SwitchElement name="isMain" label="Atur sebagai alamat utama ?" />
    </FormAddDialog>
  );
};

export default AddAddress;
