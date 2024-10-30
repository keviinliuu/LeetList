import { useState } from 'react';
import axios from 'axios';
import QuestionEntry from '../components/QuestionEntry';

interface QuestionInfo {
    title: string;
    url: string;
    difficulty: string;
}

export default function CreateList() {
    const testId = "abc";
    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [questionEntries, setQuestionEntries] = useState<QuestionInfo[]>([]);
    const [error, setError] = useState<string | null>(null);

    // Function to add a new empty QuestionEntry
    const addQuestionEntry = () => {
        setQuestionEntries([...questionEntries, { title: '', url: '', difficulty: '' }]);
    };

    // Function to update question entry data
    const updateQuestionEntry = (index: number, data: QuestionInfo) => {
        const newEntries = [...questionEntries];
        newEntries[index] = data;
        setQuestionEntries(newEntries);
    };

    // Function to handle list creation
    const handleCreateList = async () => {
        const query = `
            mutation CreateList($input: CreateListInput!) {
                createList(input: $input) {
                    ID
                    title
                    description
                    UserID
                    entries {
                        ID
                        title
                        url
                        difficulty
                        complete
                    }
                }
            }
        `;

        const variables = {
            input: {
                title,
                description,
                entries: questionEntries.map(({ title, url, difficulty }) => ({
                    title,
                    url,
                    difficulty,
                })),
                testId
            },
        };

        console.log(variables)

        try {
            const response = await axios.post('', { query, variables });
            const result = response.data;

            if (result.errors) {
                setError(result.errors[0].message);
            } else {
                console.log('List created successfully:', result.data.createList);
                // You could redirect to a success page or clear the form
            }
        } catch (err) {
            console.error('Error creating list:', err);
            setError("Failed to create the list. Please try again.");
        }
    };

    return (
        <div className='flex flex-col min-h-screen w-screen bg-richBlack text-cerulean text-3xl font-main items-center'>
            <h1 className='text-4xl p-5'>Create a New List</h1>

            {/* Title and Description Inputs */}
            <input
                type="text"
                placeholder="List Title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                className="border-2 border-gray-400 rounded-md p-2 mb-4 w-1/2 text-sm"
            />
            <textarea
                placeholder="List Description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                className="border-2 border-gray-400 rounded-md p-2 mb-4 w-1/2 text-sm"
            />

            {/* Render each QuestionEntry */}
            {questionEntries.map((_, index) => (
                <QuestionEntry
                    key={index}
                    onScrapeComplete={(data) => updateQuestionEntry(index, data)}
                />
            ))}

            {/* Plus button to add new QuestionEntry */}
            <button
                onClick={addQuestionEntry}
                className='bg-blue-500 text-white rounded-full w-10 h-10 flex items-center justify-center mt-4'
            >
                +
            </button>

            {/* Create List Button */}
            <button
                onClick={handleCreateList}
                className="bg-green-500 text-white rounded-md p-2 mt-6 text-lg"
            >
                Create List
            </button>

            {/* Error Message */}
            {error && <div className="text-red-500 mt-4">{error}</div>}
        </div>
    );
}
