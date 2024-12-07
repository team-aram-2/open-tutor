declare global {
  namespace NodeJS {
    export interface ProcessEnv {
      // Terraform //
      TERRAFORM_API_KEY: string;
    }
  }
}

export {};
