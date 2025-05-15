const TransactionCard = ({ transaction }) => {
  return (
    <div className="border rounded-lg p-4 space-y-2 shadow-sm">
      <div className="flex justify-between text-sm text-muted-foreground">
        <p className="font-medium">Payment ID: {transaction.id}</p>
        <p>{new Date(transaction.paidAt).toLocaleString()}</p>
      </div>
      <div className="text-sm">
        <p>
          <span className="font-semibold">User:</span> {transaction.fullname} (
          {transaction.userEmail})
        </p>
        <p>
          <span className="font-semibold">Order ID:</span> {transaction.orderID}
        </p>
        <p>
          <span className="font-semibold">Total:</span> Rp{" "}
          {transaction.total.toLocaleString()}
        </p>
        <p>
          <span className="font-semibold">Method:</span> {transaction.method}
        </p>
        <p>
          <span className="font-semibold">Status:</span> {transaction.status}
        </p>
      </div>
    </div>
  );
};

export default TransactionCard;
