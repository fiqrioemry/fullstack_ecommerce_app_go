const OrderCard = ({ order }) => {
  return (
    <div className="border rounded-lg p-4 space-y-2 shadow-sm">
      <div className="flex justify-between text-sm text-muted-foreground">
        <p className="font-medium">Invoice: {order.invoice}</p>
        <p>{new Date(order.createdAt).toLocaleString()}</p>
      </div>
      <div className="text-sm">
        <p>
          <span className="font-semibold">User:</span> {order.userName} (
          {order.userEmail})
        </p>
        <p>
          <span className="font-semibold">Total:</span> Rp{" "}
          {order.total.toLocaleString()}
        </p>
        <p>
          <span className="font-semibold">Status:</span> {order.status}
        </p>
      </div>
    </div>
  );
};

export default OrderCard;
