import DarkMode from "@/components/buttons/darkMode"
import LinkForm from "@/components/forms/linkForm"

export default function App() {
  return (
    <div className="h-page-screen flex flex-col">
      <div className="flex justify-end">
        <DarkMode />
      </div>
      <div className="flex flex-col justify-center items-center gap-16 flex-grow">
        <h1 className="text-4xl font-extrabold tracking-tight lg:text-5xl">ClipLink</h1>
        <LinkForm />
      </div>
    </div>
  )
}
