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

// Service untuk memperbarui profil pengguna
export const updateUserProfile = async (profileData, token) => {
  const response = await fetch(`${BASE_URL}/profile`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify(profileData),
  });
  return handleResponse(response);
};

// Service untuk membuat komunitas baru
export const createCommunity = async (communityData, token) => {
  const response = await fetch(`${BASE_URL}/communities`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify(communityData),
  });
  return handleResponse(response);
};

// --- Social Interaction Services ---

export const followUser = async (username, token) => {
  const response = await fetch(`${BASE_URL}/users/${username}/follow`, {
    method: 'POST',
    headers: { 'Authorization': `Bearer ${token}` },
  });
  return handleResponse(response);
};

export const unfollowUser = async (username, token) => {
  const response = await fetch(`${BASE_URL}/users/${username}/follow`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${token}` },
  });
  return handleResponse(response);
};

export const likePost = async (postId, token) => {
  const response = await fetch(`${BASE_URL}/posts/${postId}/like`, {
    method: 'POST',
    headers: { 'Authorization': `Bearer ${token}` },
  });
  return handleResponse(response);
};

export const unlikePost = async (postId, token) => {
  const response = await fetch(`${BASE_URL}/posts/${postId}/like`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${token}` },
  });
  return handleResponse(response);
};

// --- Community Services ---

export const getCommunityDetails = async (name) => {
  const response = await fetch(`${BASE_URL}/communities/${name}`);
  return handleResponse(response);
};

export const joinCommunity = async (name, token) => {
  const response = await fetch(`${BASE_URL}/communities/${name}/join`, {
    method: 'POST',
    headers: { 'Authorization': `Bearer ${token}` },
  });
  return handleResponse(response);
};

export const leaveCommunity = async (name, token) => {
  const response = await fetch(`${BASE_URL}/communities/${name}/join`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${token}` },
  });
  return handleResponse(response);
};