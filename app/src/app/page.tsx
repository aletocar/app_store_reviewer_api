'use client';

import { useState, useEffect } from 'react';
import AppPicker from '@/components/AppPicker';
import ReviewCard from '@/components/ReviewCard';
import { Review } from '@/types/reviews';

export default function Home() {
  const [apps, setApps] = useState<string[]>([]);
  const [selectedApp, setSelectedApp] = useState('');
  const [days, setDays] = useState(2);
  const [reviews, setReviews] = useState<Review[]>([]);

  useEffect(() => {
    fetch('http://localhost:8080/apps')
      .then(response => response.json())
      .then(data => setApps(data))
      .catch(error => console.error('Error fetching apps:', error));
  }, []);

  useEffect(() => {
    if (selectedApp && days) {
      fetch(`http://localhost:8080/apps/${selectedApp}/reviews?days=${days}`)
        .then(response => response.json())
        .then(data => {
          const sortedReviews = data.sort((a: Review, b: Review) => 
            new Date(b.updated).getTime() - new Date(a.updated).getTime()
          );
          setReviews(sortedReviews);
        })
        .catch(error => console.error('Error fetching reviews:', error));
    }
  }, [selectedApp, days]);

  return (
    <div className="min-h-screen p-8">
      <main className="max-w-4xl mx-auto">
        <h1 className="text-2xl font-bold mb-8 text-black">App Reviews</h1>

        <div className="flex gap-6 mb-8">
          <AppPicker
            apps={apps}
            selectedApp={selectedApp}
            onSelectApp={setSelectedApp}
          />

          <div className="flex flex-col gap-2">
            <label htmlFor="days" className="text-sm font-medium text-black">
              Number of days:
            </label>
            <input
              type="number"
              id="days"
              min="1"
              value={days}
              onChange={(e) =>
                setDays(Math.max(1, parseInt(e.target.value) || 1))
              }
              className="p-2 border rounded-md w-32 text-black"
            />
          </div>
        </div>

        <div className="grid gap-4">
          {reviews.map((review) => (
            <ReviewCard key={review.id} review={review} />
          ))}
        </div>
      </main>
    </div>
  );
}
