const BASE_URL = '/api';

// Fungsi helper untuk menangani response dari fetch
async function handleResponse(response) {
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.message || 'Something went wrong');
  }
  return response.json();
}

// Service untuk mengambil semua postingan
export const getPosts = async () => {
  const response = await fetch(`${BASE_URL}/posts`);
  return handleResponse(response);
};

// Service untuk membuat postingan baru
export const createPost = async (postData, token) => {
  const response = await fetch(`${BASE_URL}/posts/create`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      // Kirim token untuk autentikasi
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify(postData),
  });
  return handleResponse(response);
};

// Service untuk mengambil profil pengguna berdasarkan username
export const getUserProfile = async (username) => {
  const response = await fetch(`${BASE_URL}/users/${username}`);
  return handleResponse(response);
};

// Service untuk mengambil postingan milik seorang pengguna
export const getPostsByUsername = async (username) => {
  const response = await fetch(`${BASE_URL}/users/${username}/posts`);
  return handleResponse(response);
};