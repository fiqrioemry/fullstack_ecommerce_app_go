import { PlusIcon } from "lucide-react";
import { bannerSchema } from "@/lib/schema";
import { bannerState } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { useBannerMutation } from "@/hooks/useBanner";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { SelectElement } from "@/components/input/SelectElement";
import { InputFileElement } from "@/components/input/InputFileElement";

export const AddBanner = () => {
  const { createBanner } = useBannerMutation();

  return (
    <FormAddDialog
      state={bannerState}
      title="Create New Banner"
      schema={bannerSchema}
      loading={createBanner.isPending}
      action={createBanner.mutateAsync}
      buttonElement={
        <Button type="button">
          <PlusIcon />
          <span>Create Banner</span>
        </Button>
      }
    >
      <div className="space-y-4">
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
      </div>
    </FormAddDialog>
  );
};
