import {
  Table,
  TableRow,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { useBannersQuery } from "@/hooks/useBanner";
import { Card, CardContent } from "@/components/ui/Card";
import { SectionTitle } from "@/components/header/SectionTitle";
import { AddBanner } from "@/components/admin/banners/AddBanner";
import { DeleteBanner } from "@/components/admin/banners/DeleteBanner";
import { UpdateBanner } from "@/components/admin/banners/UpdateBanner";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";

const BannersList = () => {
  const {
    data: bannerData = [],
    isLoading,
    isError,
    refetch,
  } = useBannersQuery();

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle
        title="Banners List"
        description="Manage homepage banners based on position: Top, Side1, Side2, Bottom."
      />

      <div className="flex justify-end">
        <AddBanner />
      </div>
      <div className="rounded-md border">
        <Card className="border shadow-sm">
          <CardContent className="overflow-x-auto p-0">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="text-center w-40">Preview</TableHead>
                  <TableHead className="text-center">Position</TableHead>
                  <TableHead className="text-center w-40">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {bannerData.map((banner) => (
                  <TableRow key={banner.id}>
                    <TableCell className="flex items-center justify-center">
                      <img
                        src={banner.image}
                        alt="Banner"
                        className="h-28 w-full object-cover rounded border"
                      />
                    </TableCell>
                    <TableCell className="capitalize">
                      {banner.position}
                    </TableCell>
                    <TableCell className="space-x-4">
                      <UpdateBanner banner={banner} />
                      <DeleteBanner banner={banner} />
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    </section>
  );
};

export default BannersList;
