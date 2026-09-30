<?php

namespace App\Http\Requests;

use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Foundation\Http\FormRequest;

class StoreMatakuliahRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'program_studi_id' => ['required', 'integer', 'exists:program_studis,id'],
            'kode' => ['required', 'string', 'max:20', 'unique:matakuliahs,kode'],
            'nama' => ['required', 'string', 'max:100'],
            'sks' => ['required', 'integer', 'min:1', 'max:6'],
            'semester' => ['required', 'integer', 'min:1', 'max:8'],
        ];
    }
        public function messages(): array
    {
        return [
            'kode.unique' => 'Kode matakuliah sudah terdaftar',
            'sks.max' => 'SKS maksimal 6',
        ];
    }
}
