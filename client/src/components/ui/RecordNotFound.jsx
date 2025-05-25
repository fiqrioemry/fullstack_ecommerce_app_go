export const RecordNotFound = ({ title, q }) => {
  return (
    <div className="text-center py-16 text-muted-foreground">
      <h3 className="text-lg font-semibold">{title}</h3>
      {q && (
        <p className="text-sm mt-2">
          for keyword: <strong>{q}</strong>
        </p>
      )}
    </div>
  );
};
