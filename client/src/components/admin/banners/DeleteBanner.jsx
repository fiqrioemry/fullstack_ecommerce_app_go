import { useBannerMutation } from "@/hooks/useBanner";
import { FormDelete } from "@/components/form/FormDelete";

export const DeleteBanner = ({ banner }) => {
  const { deleteBanner } = useBannerMutation();

  const handleDeleteBanner = () => {
    deleteBanner.mutate(banner.id);
  };
  return (
    <FormDelete
      title="Delete Banners"
      loading={deleteBanner.isPending}
      onDelete={handleDeleteBanner}
      description="Are you sure want to delete this Banner?"
    />
  );
};
