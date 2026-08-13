import os

from huggingface_hub import InferenceClient


class LLMClient:

    def __init__(self):
        token = os.getenv("HF_TOKEN")

        if not token:
            raise RuntimeError(
                "HF_TOKEN is not configured"
            )

        self.client = InferenceClient(
            api_key=token,
            provider="auto",
        )

    def generate(self, prompt: str) -> str:

        response = self.client.chat.completions.create(
            model="Qwen/Qwen3-32B",
            messages=[
                {
                    "role": "user",
                    "content": prompt,
                }
            ],
        )

        content = response.choices[0].message.content

        if not content:
            raise RuntimeError(
                "LLM returned an empty response"
            )

        return content