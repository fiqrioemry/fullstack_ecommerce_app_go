import React from "react";
import { useBannersQuery } from "../../hooks/useBanner";

const BannersList = () => {
  const { data, isLoading, isError, error } = useBannersQuery(
    search,
    page,
    limit,
    sort
  );

  return (
    <section className="section space-y-6">
      <div className="text-center space-y-1 mb-6">
        <h2 className="text-2xl font-bold">Products List</h2>
        <p className="text-sm text-muted-foreground">
          See all your store product
        </p>
      </div>
    </section>
  );
};

export default BannersList;
