import { useParams } from "react-router-dom";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Loading } from "@/components/ui/Loading";
import { useProductDetailQuery } from "@/hooks/useProduct";
import { ProductInfo } from "@/components/product-detail/ProductInfo";
import { ReviewList } from "@/components/product-detail/ReviewList";
import { ProductGallery } from "@/components/product-detail/ProductGallery";

const ProductDetail = () => {
  const { slug } = useParams();
  const {
    isError,
    refetch,
    isLoading,
    data: product,
  } = useProductDetailQuery(slug);

  if (isLoading) return <Loading />;

  if (isError || !product) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section py-16 md:py-20 space-y-8">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-10">
        <ProductGallery product={product} />
        <ProductInfo product={product} />
      </div>
      <ReviewList product={product} />
    </section>
  );
};

export default ProductDetail;
