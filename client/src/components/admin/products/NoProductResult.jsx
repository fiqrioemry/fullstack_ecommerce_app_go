const NoProductResult = ({ search }) => {
  return (
    <div className="text-center py-16 text-muted-foreground">
      <h3 className="text-lg font-semibold">No products found</h3>
      {search && (
        <p className="text-sm mt-2">
          for keyword: <strong>{search}</strong>
        </p>
      )}
    </div>
  );
};

export default NoProductResult;
