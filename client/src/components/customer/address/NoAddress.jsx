export const NoAddress = () => (
  <div className="border border-dashed border-border bg-muted/40 rounded-xl p-10 text-center space-y-4">
    <div className="flex justify-center">
      <img src="/no-transactions.webp" alt="no-transactions" className="h-72" />
    </div>

    <h3 className="text-lg font-semibold text-foreground">No Address Found</h3>
  </div>
);
