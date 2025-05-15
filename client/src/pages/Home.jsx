import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
import { Swiper, SwiperSlide } from "swiper/react";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Loading } from "@/components/ui/Loading";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { useCategoriesQuery } from "@/hooks/useCategory";
import { Autoplay, Pagination, Navigation } from "swiper/modules";
import { ProductCard } from "@/components/product-results/ProductCard";
import { useBannersQuery } from "@/hooks/useBanner";

const Home = () => {
  const { data: bannerData } = useBannersQuery();
  const { data: categoryData } = useCategoriesQuery();
  const {
    data: productData,
    isError,
    refetch,
    isLoading,
  } = useSearchProductsQuery({ limit: 50 });

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const banners = bannerData || [];
  const categories = categoryData || [];
  const products = productData?.data || [];

  const topBanners = banners.filter((b) => b.position === "top");
  const side1Banners = banners.filter((b) => b.position === "side1");
  const side2Banners = banners.filter((b) => b.position === "side2");
  const bottomBanners = banners.filter((b) => b.position === "bottom");

  const featuredProducts = products.filter((p) => p.isFeatured);
  const discountProducts = products.filter((p) => p.discount > 0.05);
  const gadgetProducts = products.filter(
    (p) => p.category === "Gadget & Electronics"
  );
  console.log(productData);

  return (
    <section className="section py-16 md:py-20 space-y-8">
      {topBanners.length > 0 && (
        <div className="mb-6">
          <Swiper
            navigation
            centeredSlides
            spaceBetween={10}
            loop={topBanners.length > 1}
            pagination={{ clickable: true }}
            modules={[Autoplay, Pagination, Navigation]}
            autoplay={{ delay: 4000, disableOnInteraction: false }}
            className="rounded-xl overflow-hidden shadow-lg"
          >
            {topBanners.map((banner) => (
              <SwiperSlide key={banner.id}>
                <img
                  src={banner.imageUrl}
                  alt="top-banner"
                  className="h-64 md:h-96 w-full object-cover"
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      )}
      {/* Browse by Category */}
      {categories.length > 0 && (
        <section className="mb-8">
          <h2 className="text-lg font-semibold mb-4">🛍️ Browse by Category</h2>
          <div className="flex gap-4 overflow-x-auto pb-2 no-scrollbar">
            {categories.map((cat) => (
              <div
                key={cat.id}
                onClick={() =>
                  (window.location.href = `/products?category=${cat.slug}`)
                }
                className="min-w-[140px] cursor-pointer rounded-xl border shadow hover:shadow-md transition bg-background"
              >
                <img
                  src={cat.imageUrl}
                  alt={cat.name}
                  className="object-cover rounded-t-xl"
                />
                <div className="p-2 text-center text-sm font-medium">
                  {cat.name}
                </div>
              </div>
            ))}
          </div>
        </section>
      )}

      {side1Banners.length > 0 && (
        <div className="md:hidden mb-6">
          <Swiper
            loop
            spaceBetween={10}
            autoplay={{ delay: 4000 }}
            pagination={{ clickable: true }}
            modules={[Autoplay, Pagination]}
            className="rounded-xl overflow-hidden shadow-md"
          >
            {side1Banners.map((banner) => (
              <SwiperSlide key={banner.id}>
                <img
                  src={banner.imageUrl}
                  alt="side1-mobile"
                  className="h-48 w-full object-cover"
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      )}

      {/* Grid for Side1 - Products - Side2 */}
      <div className="grid grid-cols-1 md:grid-cols-12 gap-6 mb-8">
        {side2Banners.length > 0 && (
          <div className="hidden md:block md:col-span-2 space-y-4">
            {side2Banners.map((banner) => (
              <img
                key={banner.id}
                src={banner.imageUrl}
                alt="side2"
                className="rounded-xl object-cover w-full shadow"
              />
            ))}
          </div>
        )}
        {/* Main Content: Featured & Discount */}
        <div className="md:col-span-8 space-y-8">
          {/* Featured Products */}
          <section>
            <h2 className="text-xl font-semibold mb-2">✨ Featured Products</h2>
            <p className="text-sm text-muted-foreground mb-4">
              Do not miss the featured deals! It Limited time sales.
            </p>
            <Swiper
              spaceBetween={16}
              navigation
              modules={[Navigation]}
              breakpoints={{
                0: { slidesPerView: 1 },
                640: { slidesPerView: 2 },
                768: { slidesPerView: 3 },
                1024: { slidesPerView: 4 },
              }}
            >
              {featuredProducts.map((product) => (
                <SwiperSlide key={product.id}>
                  <ProductCard product={product} />
                </SwiperSlide>
              ))}
            </Swiper>
          </section>

          {/* Discount Products */}
          <section>
            <h2 className="text-xl font-semibold mb-2">
              🔥 Discounted Products
            </h2>
            <p className="text-sm text-muted-foreground mb-4">
              Limited time discounts! Get our product with most valuable price.
            </p>
            <Swiper
              spaceBetween={16}
              navigation
              modules={[Navigation]}
              breakpoints={{
                0: { slidesPerView: 1 },
                640: { slidesPerView: 2 },
                768: { slidesPerView: 3 },
                1024: { slidesPerView: 4 },
              }}
            >
              {discountProducts.map((product) => (
                <SwiperSlide key={product.id}>
                  <ProductCard product={product} />
                </SwiperSlide>
              ))}
            </Swiper>
          </section>
        </div>

        {side1Banners.length > 0 && (
          <div className="hidden md:block md:col-span-2 space-y-4">
            {side1Banners.map((banner) => (
              <img
                key={banner.id}
                src={banner.imageUrl}
                alt="side1"
                className="rounded-xl object-cover w-full shadow"
              />
            ))}
          </div>
        )}
      </div>

      {side2Banners.length > 0 && (
        <div className="md:hidden mb-6">
          <Swiper
            loop
            spaceBetween={10}
            autoplay={{ delay: 4000 }}
            pagination={{ clickable: true }}
            modules={[Autoplay, Pagination]}
            className="rounded-xl overflow-hidden shadow-md"
          >
            {side2Banners.map((banner) => (
              <SwiperSlide key={banner.id}>
                <img
                  src={banner.imageUrl}
                  alt="side2-mobile"
                  className="h-48 w-full object-cover"
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      )}

      {/* Gadget & Electronics */}
      <section className="mb-8">
        <h2 className="text-xl font-semibold mb-2">📱 Gadget & Electronics</h2>
        <p className="text-sm text-muted-foreground mb-4">
          Discover the latest tech at great prices.
        </p>
        <Swiper
          spaceBetween={16}
          navigation
          modules={[Navigation]}
          breakpoints={{
            0: { slidesPerView: 1 },
            640: { slidesPerView: 2 },
            768: { slidesPerView: 3 },
            1024: { slidesPerView: 4 },
          }}
        >
          {gadgetProducts.map((product) => (
            <SwiperSlide key={product.id}>
              <ProductCard product={product} />
            </SwiperSlide>
          ))}
        </Swiper>
      </section>

      {bottomBanners.length > 0 && (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {bottomBanners.map((banner) => (
            <img
              key={banner.id}
              src={banner.imageUrl}
              alt="bottom-banner"
              className="rounded-xl object-cover h-56 md:h-64 w-full shadow"
            />
          ))}
        </div>
      )}
    </section>
  );
};

export default Home;
