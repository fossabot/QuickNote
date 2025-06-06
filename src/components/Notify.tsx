import { toast, Toaster } from "react-hot-toast";

export function Notify() {
  // 触发通知的函数
  const showNotification = () => {
    console.log("Success!");
    toast.success("Success!");
  };

  return (
    <div>
      <button onClick={showNotification}>显示通知</button>
      <Toaster position="top-right" />
    </div>
  );
}
