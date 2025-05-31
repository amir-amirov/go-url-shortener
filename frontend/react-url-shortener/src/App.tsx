import { useState } from 'react';
import './App.css';

function App() {
  const [rawURL, setRawURL] = useState('');
  const [shortURL, setShortURL] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    setError('');
    setShortURL('');
    setLoading(true);

    try {
      // const response = await fetch('http://localhost:8080/shorten', {
      const response = await fetch('/shorten', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url: rawURL }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Something went wrong');
      }

      const data = await response.json();
      setShortURL(data.short_url);
    } catch (err: any) {
      setError(err.message || 'Something went wrong. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <div className="card">
        <h1>URL Shortener</h1>
        <form onSubmit={handleSubmit}>
          <label htmlFor="url">Enter URL to shorten</label>
          <input
            type="url"
            id="url"
            value={rawURL}
            onChange={(e) => setRawURL(e.target.value)}
            placeholder="https://example.com"
            required
          />
          <button type="submit" disabled={loading}>
            {loading ? 'Shortening...' : 'Shorten URL'}
          </button>
        </form>
        {shortURL && (
          <div className="success">
            <p>
              Shortened URL:{' '}
              <a href={shortURL} target="_blank" rel="noopener noreferrer">
                {shortURL}
              </a>
            </p>
          </div>
        )}
        {error && (
          <div className="error">
            <p>{error}</p>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;