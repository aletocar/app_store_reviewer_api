interface AppPickerProps {
  apps: string[];
  selectedApp: string;
  onSelectApp: (app: string) => void;
}

export default function AppPicker({ apps, selectedApp, onSelectApp }: AppPickerProps) {
  return (
    <div className="flex flex-col gap-2">
      <label htmlFor="app-select" className="text-sm font-medium text-black">
        Select an app:
      </label>
      <select
        id="app-select"
        value={selectedApp}
        onChange={(e) => onSelectApp(e.target.value)}
        className="p-2 border rounded-md bg-white text-black"
      >
        <option value="">Select an app...</option>
        {apps.map(app => (
          <option key={app} value={app}>
            {app}
          </option>
        ))}
      </select>
    </div>
  );
} 