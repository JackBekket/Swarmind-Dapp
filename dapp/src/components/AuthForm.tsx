import React, { useState } from 'react';

interface Props {
  onSubmit: (email: string, password: string) => void;
}

const AuthForm: React.FC<Props> = ({ onSubmit }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(email, password);
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="w-[480px] bg-white border border-[#F4F4F6] shadow-md rounded-3xl px-8 py-12 flex flex-col items-center gap-8"
    >
      <h1 className="font-[Coolvetica] text-[64px] leading-[48px] tracking-[-0.01em] text-[#5E64FF]">Swarmind</h1>

      <div className="w-full flex flex-col gap-4">
        <label className="text-sm font-medium text-[#8C929C] flex flex-col gap-1">
          Email
          <input
            type="email"
            className="w-full h-12 px-3 rounded-xl bg-[#F4F4F6] focus:outline-none"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </label>

        <label className="text-sm font-medium text-[#8C929C] flex flex-col gap-1">
          Password
          <div className="relative w-full">
            <input
              type={showPassword ? 'text' : 'password'}
              className="w-full h-12 px-3 rounded-xl bg-[#F4F4F6] focus:outline-none pr-10"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              minLength={8}
            />
            <button
              type="button"
              onClick={() => setShowPassword((prev) => !prev)}
              className="absolute right-3 top-1/2 -translate-y-1/2 text-xs text-[#5B616D]"
            >
              {showPassword ? 'Hide' : 'Show'}
            </button>
          </div>
        </label>
      </div>

      <button
        type="submit"
        className="w-full h-12 bg-[#5E64FF] text-white font-semibold rounded-xl shadow-md flex items-center justify-center gap-1"
      >
        Sign in
      </button>
    </form>
  );
};

export default AuthForm; 