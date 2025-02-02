'use client';

import { useState } from 'react';
import { Review } from '@/types/reviews';

interface ReviewCardProps {
  review: Review;
}

export default function ReviewCard({ review }: ReviewCardProps) {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const maxContentLength = 200;
  const isContentTruncated = review.content.length > maxContentLength;
  const displayContent = isContentTruncated 
    ? `${review.content.substring(0, maxContentLength)}...` 
    : review.content;

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleString();
  };

  return (
    <div className="p-4 border rounded-lg shadow-sm bg-white">
      <div className="flex justify-between items-start mb-2">
        <div className="text-yellow-400">{'⭐'.repeat(parseInt(review.rating))}</div>
        <span className="text-sm text-black">
          {formatDate(review.updated)}
        </span>
      </div>
      <div className="mb-2">
        <span className="text-sm text-black">{review.author}</span>
      </div>
      <h3 className="font-bold mb-2 text-black">{review.title}</h3>
      <p className="text-sm text-black">
        {displayContent}
        {isContentTruncated && (
          <button
            onClick={() => setIsModalOpen(true)}
            className="ml-2 text-blue-500 hover:underline"
          >
            more
          </button>
        )}
      </p>

      {isModalOpen && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
          <div className="bg-white p-6 rounded-lg max-w-2xl w-full max-h-[90vh] overflow-y-auto">
            <h3 className="font-bold mb-4">{review.title}</h3>
            <p className="text-black mb-4">{review.content}</p>
            <button
              onClick={() => setIsModalOpen(false)}
              className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              Close
            </button>
          </div>
        </div>
      )}
    </div>
  );
} 