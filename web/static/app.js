const api = {
  async request(path, options = {}) {
    const headers = {'Content-Type': 'application/json', ...(options.headers || {})};
    const token = localStorage.getItem('checkpoint_token');
    if (token) headers.Authorization = `Bearer ${token}`;
    const response = await fetch(path, {...options, headers});
    const body = await response.json().catch(() => ({}));
    if (!response.ok) throw new Error(body.error || 'La API no pudo completar la acción');
    return body.data;
  }
};

function authPage(mode) {
  return { form: {username: '', display_name: '', password: ''}, loading: false, error: '', init() {}, async submit() {
    this.loading = true; this.error = '';
    try {
      if (mode === 'register') await api.request('/users', {method: 'POST', body: JSON.stringify(this.form)});
      const result = await api.request('/auth/login', {method: 'POST', body: JSON.stringify({username: this.form.username, password: this.form.password})});
      localStorage.setItem('checkpoint_token', result.token); window.location.href = '/app';
    } catch (error) { this.error = error.message; } finally { this.loading = false; }
  }};
}

function communityApp() {
  return {view: 'feed', posts: [], rooms: [], user: null, newPost: '', newRoom: '', error: '', async init() {
    if (!localStorage.getItem('checkpoint_token')) { window.location.href = '/login'; return; }
    await this.loadPosts();
  }, async loadPosts() { try { this.posts = await api.request('/posts') || []; } catch (error) { this.error = error.message; } }, async loadRooms() { try { this.rooms = await api.request('/rooms') || []; } catch (error) { this.error = error.message; } }, async createPost() { try { await api.request('/posts', {method:'POST', body:JSON.stringify({content:this.newPost})}); this.newPost = ''; await this.loadPosts(); } catch (error) { this.error = error.message; } }, async createRoom() { if (!this.newRoom.trim()) return; try { await api.request('/rooms', {method:'POST', body:JSON.stringify({name:this.newRoom})}); this.newRoom = ''; await this.loadRooms(); } catch (error) { this.error = error.message; } }, async like(post) { try { await api.request(`/posts/${post.id}/likes`, {method:'POST'}); post.likes++; } catch (error) { this.error = error.message; } }, async openComments(post) { try { post.comments = await api.request(`/posts/${post.id}/comments`) || []; } catch (error) { this.error = error.message; } }, async joinRoom(room) { window.location.href = `/chat/${room.id}`; }, logout() { localStorage.removeItem('checkpoint_token'); window.location.href = '/login'; }};
}
