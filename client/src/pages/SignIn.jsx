import { Link } from "react-router-dom";
import { loginSchema } from "@/lib/schema";
import { getLoginState } from "@/lib/constant";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { FormInput } from "@/components/form/FormInput";
import { GoogleOAuth } from "@/components/public/GoogleOAuth";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputTextElement } from "@/components/input/InputTextElement";

const SignIn = () => {
  const { login, loading, rememberMe } = useAuthStore();

  return (
    <section className="min-h-screen flex items-center justify-center bg-background px-4">
      <div className="w-full max-w-4xl grid grid-cols-1 md:grid-cols-2 bg-card border border-border rounded-xl shadow-lg overflow-hidden">
        {/* Left Side (Illustration) */}
        <div className="hidden md:block bg-blue-600 p-8 text-white text-center">
          <h2 className="text-3xl font-bold mb-4">Welcome Back!</h2>
          <p className="text-sm">Login and explore your dashboard</p>
          <img
            src="/signin-wallpaper.webp"
            alt="sign-in-illustration"
            className="mt-6 w-full h-auto"
          />
        </div>

        <div className="px-6 py-10">
          <div className="mb-6 flex justify-center text-center">
            <WebLogo />
          </div>

          <FormInput
            action={login}
            text="Sign In"
            className="w-full"
            schema={loginSchema}
            isLoading={loading}
            state={getLoginState(rememberMe)}
          >
            <InputTextElement
              name="email"
              label="Email"
              placeholder="Enter your email"
            />
            <InputTextElement
              name="password"
              label="Password"
              type="password"
              placeholder="********"
            />
            <SwitchElement name="rememberMe" label="Remember Me" />
          </FormInput>
          <div className="text-center py-2 text-sm text-muted-foreground">
            Or
          </div>

          <GoogleOAuth buttonText="Sign in with google" />
          <p className="text-sm text-center mt-6 text-muted-foreground">
            Don't have an account?{" "}
            <Link
              to="/signup"
              className="text-primary font-medium hover:underline"
            >
              Sign up now
            </Link>
          </p>
        </div>
      </div>
    </section>
  );
};

export default SignIn;
