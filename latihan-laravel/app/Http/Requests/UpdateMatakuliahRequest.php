<?php

namespace App\Http\Requests;

use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Foundation\Http\FormRequest;

class UpdateMatakuliahRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return false;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        $id = $this->route('matakuliah')->id;

        return [
            'program_studi_id' => ['sometimes', 'integer', 'exists:program_studis,id'],
            'kode' => ['sometimes', 'string', 'max:20', 'unique:matakuliahs,kode,' . $id],
            'nama' => ['sometimes', 'string', 'max:100'],
            'sks' => ['sometimes', 'integer', 'min:1', 'max:6'],
            'semester' => ['sometimes', 'integer', 'min:1', 'max:8'],
        ];
    }
}
