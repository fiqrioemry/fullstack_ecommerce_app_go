import { bannerSchema } from "@/lib/schema";
import { useBannerMutation } from "@/hooks/useBanner";
import { SelectElement } from "@/components/input/SelectElement";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputFileElement } from "@/components/input/InputFileElement";

export const UpdateBanner = ({ banner }) => {
  const { updateBanner } = useBannerMutation();

  return (
    <FormUpdateDialog
      state={banner}
      title="Update banner"
      schema={bannerSchema}
      loading={updateBanner.isPending}
      action={updateBanner.mutateAsync}
    >
      <SelectElement
        name="position"
        label="Position"
        placeholder="Select position for the banner"
        options={["top", "side-right", "side-left", "bottom"]}
      />

      <InputFileElement
        isSingle
        name="image"
        label="Banner Image"
        placeholder="Upload banner image"
        note="You can only upload .jpg, .jpeg, .png files"
      />
    </FormUpdateDialog>
  );
};
