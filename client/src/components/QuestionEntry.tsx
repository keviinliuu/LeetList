import { useState } from 'react';
import axios from 'axios';

interface QuestionInfo {
    title: string;
    url: string;
    difficulty: string;
}

interface QuestionEntryProps {
    onScrapeComplete: (data: QuestionInfo) => void;
}

export default function QuestionEntry({ onScrapeComplete }: QuestionEntryProps) {
    const [questionURL, setQuestionURL] = useState<string>('');
    const [questionInfo, setQuestionInfo] = useState<QuestionInfo | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    const handlePaste = async (e: React.ClipboardEvent<HTMLInputElement>) => {
        e.preventDefault();
        const pasteText = e.clipboardData.getData('text');
        setQuestionURL(pasteText);
        await handleScrapeQuestion(pasteText);
    };

    const handleScrapeQuestion = async (url: string) => {
        setLoading(true);
        const query = `
            query ScrapeQuestion($url: String!) {
                scrapeQuestion(url: $url) {
                    title
                    difficulty
                }
            }
        `;

        const variables = { url };

        try {
            const response = await axios.post('', { query, variables });
            const result = response.data;

            if (result.errors) {
                setError(result.errors[0].message);
            } else {
                const questionData = result.data.scrapeQuestion;
                setQuestionInfo(questionData);
                onScrapeComplete({ title: questionData.title, url, difficulty: questionData.difficulty });
            }
        } catch (err) {
            console.error('Scrape error:', err);
            setError("Failed to retrieve question info. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="p-2">
            {loading ? (
                <div className="text-gray-500 text-sm">Loading...</div>
            ) : questionInfo ? (
                <div className="text-white p-2 bg-gray-800 rounded-md">
                    <div className="text-lg font-semibold">{questionInfo.title}</div>
                    <div className="text-sm text-gray-400">Difficulty: {questionInfo.difficulty}</div>
                </div>
            ) : (
                <input
                    type="text"
                    className="border-2 border-gray-400 rounded-md p-2 w-full text-sm"
                    placeholder="Paste LeetCode question URL here"
                    value={questionURL}
                    onPaste={handlePaste}
                />
            )}
            {error && <div className="text-red-500 mt-2">{error}</div>}
        </div>
    );
}
