export const sendPasswordResetEmail = async (email: string): Promise<void> => {
  // Placeholder implementation simulating async network request
  console.info('[stub] sendPasswordResetEmail', email);
  return new Promise((resolve) => setTimeout(resolve, 500));
};

export const confirmPasswordReset = async (
  oobCode: string,
  newPassword: string,
): Promise<void> => {
  console.info('[stub] confirmPasswordReset', oobCode, newPassword);
  return new Promise((resolve) => setTimeout(resolve, 500));
}; 